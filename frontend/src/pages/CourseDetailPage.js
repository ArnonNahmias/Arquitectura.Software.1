import React, { useEffect, useState, useContext } from 'react';
import { useParams } from 'react-router-dom';
import axios from 'axios';
import { Box, Typography, CircularProgress, Alert, Card, CardMedia, CardContent, Button } from '@mui/material';
import AuthContext from '../context/AuthContext';

const CourseDetailPage = () => {
  const { id } = useParams();
  const [course, setCourse] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [subscribeError, setSubscribeError] = useState(null);
  const [subscribeSuccess, setSubscribeSuccess] = useState(null);
  const { user } = useContext(AuthContext);

  useEffect(() => {
    const fetchCourse = async () => {
      try {
        const response = await axios.get(`http://localhost:8080/courses/${id}`);
        setCourse(response.data);
      } catch (error) {
        console.error('Error fetching course details', error);
        setError('Error fetching course details');
      } finally {
        setLoading(false);
      }
    };

    fetchCourse();
  }, [id]);

  const handleSubscribe = async () => {
    console.log('User:', user.userId); // Add this line for debugging
    if (!user || !user.userId) {
      setSubscribeError('You must be logged in to subscribe');
      return;
    }
  
    try {
      console.log(parseInt(user.userId), "hola")
      const response = await axios.post('http://localhost:8080/subscriptions', {
        userID: parseInt(user.userId),
        courseID: parseInt(id)
      });
      console.log('Subscribed successfully:', response.data);
      setSubscribeSuccess('Subscribed successfully!');
    } catch (err) {
      console.error('Error subscribing to the course', err);
      setSubscribeError('Error subscribing to the course');
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

  return (
    <Box sx={{ flexGrow: 1, p: 3 }}>
      <Card>
        {course?.imageURL && (
          <Box sx={{ maxWidth: 500, maxHeight: 500, mx: 'auto' }}>
            <CardMedia
              component="img"
              image={course.imageURL}
              alt={course.nombre}
              sx={{ maxWidth: 500, maxHeight: 500 }}
            />
          </Box>
        )}
        <CardContent>
          <Typography variant="h4" component="h1" gutterBottom>
            {course.nombre}
          </Typography>
          <Button
            variant="contained"
            color="primary"
            onClick={handleSubscribe}
            sx={{ mb: 2 }}
          >
            Subscribe to this Course
          </Button>
          {subscribeError && <Alert severity="error">{subscribeError}</Alert>}
          {subscribeSuccess && <Alert severity="success">{subscribeSuccess}</Alert>}
          <Typography variant="body1" paragraph>
            {course.descripcion}
          </Typography>
          <Typography variant="body1" paragraph>
            Category: {course.categoria}
          </Typography>
          <Typography variant="body1" paragraph>
            Difficulty: {course.dificultad}
          </Typography>
          <Typography variant="body1" paragraph>
            Price: ${course.precio}
          </Typography>
        </CardContent>
      </Card>
    </Box>
  );
};

export default CourseDetailPage;
