import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { fetchCourses } from '../services/api';
import { Box, Card, CardContent, CardMedia, Typography, Grid, CircularProgress, Alert, TextField, Button } from '@mui/material';

const HomePage = () => {
  const [courses, setCourses] = useState([]);
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(true);
  const [searchQuery, setSearchQuery] = useState('');
  const [searchResults, setSearchResults] = useState([]);

  useEffect(() => {
    const getCourses = async () => {
      try {
        const coursesData = await fetchCourses();
        setCourses(coursesData);
      } catch (error) {
        console.error('Error fetching courses', error);
        setError('Error fetching courses');
      } finally {
        setLoading(false);
      }
    };

    getCourses();
  }, []);

  const handleSearch = async () => {
    if (searchQuery.trim() === '') {
      setSearchResults([]);
      return;
    }

    setLoading(true);
    setError(null);

    try {
      const responseById = await axios.get(`http://localhost:8080/courses/${searchQuery}`);
      const responseByName = await axios.get(`http://localhost:8080/courses/name/${searchQuery}`);
      
      const coursesById = responseById.data ? [responseById.data] : [];
      const coursesByName = responseByName.data || [];
      
      const combinedResults = [...coursesById, ...coursesByName];
      
      setSearchResults(combinedResults);
    } catch (error) {
      console.error('Error fetching courses:', error.response?.data || error.message);
      setError(`Error fetching courses: ${error.response?.data?.error || error.message}`);
    } finally {
      setLoading(false);
    }
  };

  if (loading) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
        <CircularProgress />
      </Box>
    );
  }

  if (error) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
        <Alert severity="error">{error}</Alert>
      </Box>
    );
  }

  const displayedCourses = searchResults.length > 0 ? searchResults : courses;

  return (
    <Box sx={{ flexGrow: 1, p: 3 }}>
      <Typography variant="h4" component="h1" gutterBottom>
        Available Courses
      </Typography>
      <Box display="flex" alignItems="center" mb={2}>
        <TextField
          label="Search Courses"
          value={searchQuery}
          onChange={(e) => setSearchQuery(e.target.value)}
          variant="outlined"
          fullWidth
          margin="normal"
        />
        <Button variant="contained" color="primary" onClick={handleSearch} style={{ marginLeft: '16px', height: '56px' }}>
          Search
        </Button>
      </Box>
      <Grid container spacing={3}>
        {displayedCourses.map(course => (
          <Grid item xs={12} sm={6} md={4} key={course.ID}>
            <Card>
              <CardMedia
                component="img"
                height="auto"
                image={course.imageURL}
                alt={course.nombre}
              />
              <CardContent>
                <Typography gutterBottom variant="h5" component="div">
                  {course.nombre}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Difficulty: {course.dificultad}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Price: ${course.precio}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Description: {course.descripcion}
                </Typography>
              </CardContent>
            </Card>
          </Grid>
        ))}
      </Grid>
    </Box>
  );
};

export default HomePage;
