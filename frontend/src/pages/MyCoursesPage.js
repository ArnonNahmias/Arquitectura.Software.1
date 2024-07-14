import React, { useEffect, useState, useContext } from 'react';
import { fetchUserSubscriptions } from '../services/api';
import AuthContext from '../context/AuthContext';
import { Container, Typography, Grid, Card, CardContent, CardMedia, CircularProgress, Alert } from '@mui/material';

const MyCoursesPage = () => {
    const [courses, setCourses] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const { user } = useContext(AuthContext);

    useEffect(() => {
        const fetchSubscriptions = async () => {
            try {
                console.log(`Fetching subscriptions for user ID: ${user.userId}`);
                const data = await fetchUserSubscriptions(user.userId);
                console.log('Fetched courses:', data); // Log the data to inspect it
                setCourses(data);
            } catch (error) {
                console.error('Error fetching subscriptions:', error);
                setError('Failed to load subscriptions. Please try again later.');
            } finally {
                setLoading(false);
            }
        };

        if (user) {
            fetchSubscriptions();
        }
    }, [user]);

    if (loading) {
        return (
            <Container>
                <Typography variant="h4" gutterBottom>My Courses</Typography>
                <CircularProgress />
            </Container>
        );
    }

    if (error) {
        return (
            <Container>
                <Typography variant="h4" gutterBottom>My Courses</Typography>
                <Alert severity="error">{error}</Alert>
            </Container>
        );
    }

    return (
        <Container sx={{ padding: 3 }}>
            <Typography variant="h4" gutterBottom>My Courses</Typography>
            <Grid container spacing={3}>
                {courses.map((course, index) => (
                    <Grid item xs={12} sm={6} md={4} key={`${course.ID}-${index}`}>
                        <Card sx={{ marginBottom: 2 }}>
                            <CardMedia
                                component="img"
                                height="140"
                                image={course.imageURL || course.ImageURL} // Check for both possible field names
                                alt={course.nombre}
                            />
                            <CardContent>
                                <Typography variant="h6">{course.nombre}</Typography>
                                <Typography variant="body2" color="textSecondary">{course.descripcion}</Typography>
                                <Typography variant="body2" color="textSecondary">Category: {course.categoria}</Typography>
                                <Typography variant="body2" color="textSecondary">Price: ${course.precio}</Typography>
                            </CardContent>
                        </Card>
                    </Grid>
                ))}
            </Grid>
        </Container>
    );
};

export default MyCoursesPage;
