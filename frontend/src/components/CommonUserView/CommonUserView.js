// src/components/CommonUserView.js
import React, { useContext, useState } from 'react';
import { Container, Row, Col, Card, Button, Alert, Form } from 'react-bootstrap';
import { CourseContext } from '../../contexts/CourseContext';
import './CommonUserView.scss';

const CommonUserView = () => {
  const { courses, enrolledCourses, enrollCourse } = useContext(CourseContext);
  const [message, setMessage] = useState(null);
  const [searchTerm, setSearchTerm] = useState('');

  const handleEnroll = (course) => {
    if (enrolledCourses.some(c => c.title === course.title)) {
      setMessage(`Ya estás inscrito en el curso: ${course.title}`);
    } else {
      enrollCourse(course);
      setMessage(`Felicitaciones por inscribirte al curso: ${course.title}`);
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
        placeholder="Buscar cursos"
        value={searchTerm}
        onChange={handleSearchChange}
        className="mb-4"
      />
      <Row>
        {filteredCourses.map((course, index) => (
          <Col key={index} xs={12} md={6} lg={4} className="course-col">
            <Card className="course-card">
              {course.imageUrl && <Card.Img variant="top" src={course.imageUrl} alt={course.title} />}
              <Card.Body>
                <Card.Title>{course.title || 'No Title'}</Card.Title>
                <Card.Text>{course.description || 'No Description'}</Card.Text>
                <Card.Text><strong>Price:</strong> ${course.price}</Card.Text>
                <Button variant="primary" onClick={() => handleEnroll(course)}>Inscribirse</Button>
              </Card.Body>
            </Card>
          </Col>
        ))}
      </Row>
    </Container>
  );
};

export default CommonUserView;
