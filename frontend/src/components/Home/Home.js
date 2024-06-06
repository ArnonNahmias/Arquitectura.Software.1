// src/components/Home.js
import React, { useContext, useEffect, useState } from 'react';
import { Container, Row, Col, Card, Button, Alert, Form } from 'react-bootstrap';
import { CourseContext } from '../../contexts/CourseContext';
import { useNavigate } from 'react-router-dom';
import './Home.scss';

const CommonUserView = ({ userRole }) => {
  const { courses, enrolledCourses, enrollCourse } = useContext(CourseContext);
  const [cursos, setCursos] = useState([]);
  const [message, setMessage] = useState(null);
  const [searchTerm, setSearchTerm] = useState('');
  const navigate = useNavigate();

  useEffect(() => {
    console.log(`Fetching data from http://localhost:8080/search?query=${searchTerm}`)
    // Fetch data from the API based on the search term
    fetch(`http://localhost:8080/search?query=${searchTerm}`)
      .then(response => response.json())
      .then(data => setCursos(data.results))
      .catch(error => console.error('Error fetching courses:', error));
  }, [searchTerm]);

  const handleEnroll = (course) => {
    if (userRole !== 'commonUser') {
      setMessage('You must be logged in as a common user to subscribe to courses.');
      setTimeout(() => {
        navigate('/login');
      }, 1500); // Redirect after showing the message for 1.5 seconds
      return;
    }

    if (enrolledCourses.some(c => c.title === course.title)) {
      setMessage(`You are already enrolled in the course: ${course.title}`);
    } else {
      enrollCourse(course);
      setMessage(`Congratulations on enrolling in the course: ${course.title}`);
    }
    setTimeout(() => setMessage(null), 3000); // Clear message after 3 seconds
  };

  const handleSearchChange = (e) => {
    setSearchTerm(e.target.value);
  };

  const filteredCourses = courses.filter(course =>
    course.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
    course.description.toLowerCase().includes(searchTerm.toLowerCase())
  );

  return (
    <Container className="common-user-view">
      {message && <Alert variant="success">{message}</Alert>}
      <Form.Control
        type="text"
        placeholder="Search courses"
        value={searchTerm}
        onChange={handleSearchChange}
        className="mb-4"
      />
      <Row>
        {cursos.map((course, index) => (
          <Col key={index} xs={12} md={6} lg={4} className="course-col">
            <Card className="course-card">
              {course.imageUrl && <Card.Img variant="top" src={course.imageUrl} alt={course.title} />}
              <Card.Body>
                <Card.Title>{course.title || 'No Title'}</Card.Title>
                <Card.Text>{course.description || 'No Description'}</Card.Text>
                <img src={course.image_url} />
                <Button variant="primary" onClick={() => handleEnroll(course)}>Enroll</Button>
              </Card.Body>
            </Card>
          </Col>
        ))}
      </Row>
    </Container>
  );
};

export default CommonUserView;
