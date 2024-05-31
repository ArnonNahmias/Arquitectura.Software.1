import React, { useState } from 'react';
import { Container, Row, Col, Card, Form, Button } from 'react-bootstrap';
import './Courses.scss';

const CoursePage = () => {
  const [courses, setCourses] = useState([]);
  const [courseInput, setCourseInput] = useState({
    title: '',
    description: '',
    imageFile: null,
    imageUrl: ''
  });

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setCourseInput((prevInput) => ({ ...prevInput, [name]: value }));
  };

  const handleFileChange = (e) => {
    const file = e.target.files[0];
    const reader = new FileReader();
    reader.onloadend = () => {
      setCourseInput((prevInput) => ({ ...prevInput, imageFile: file, imageUrl: reader.result }));
    };
    if (file) {
      reader.readAsDataURL(file);
    }
  };

  const handleAddCourse = () => {
    setCourses((prevCourses) => [...prevCourses, courseInput]);
    setCourseInput({ title: '', description: '', imageFile: null, imageUrl: '' });
  };

  return (
    <Container className="course-page">
      <Row>
        <Col xs={12}>
          <h2>Add a New Course</h2>
          <Form>
            <Form.Group controlId="courseTitle">
              <Form.Label>Title</Form.Label>
              <Form.Control
                type="text"
                placeholder="Enter course title"
                name="title"
                value={courseInput.title}
                onChange={handleInputChange}
              />
            </Form.Group>
            <Form.Group controlId="courseDescription">
              <Form.Label>Description</Form.Label>
              <Form.Control
                type="text"
                placeholder="Enter course description"
                name="description"
                value={courseInput.description}
                onChange={handleInputChange}
              />
            </Form.Group>
            <Form.Group controlId="courseImageFile">
              <Form.Label>Image File</Form.Label>
              <Form.Control
                type="file"
                name="imageFile"
                onChange={handleFileChange}
              />
            </Form.Group>
            <Button variant="primary" onClick={handleAddCourse}>
              Add Course
            </Button>
          </Form>
        </Col>
      </Row>
      <Row className="mt-4">
        {courses.map((course, index) => (
          <Col key={index} xs={12} md={6} lg={4} className="course-col">
            <Card className="course-card">
              <Card.Img variant="top" src={course.imageUrl} alt={course.title} />
              <Card.Body>
                <Card.Title>{course.title}</Card.Title>
                <Card.Text>{course.description}</Card.Text>
              </Card.Body>
            </Card>
          </Col>
        ))}
      </Row>
    </Container>
  );
};

export default CoursePage;
