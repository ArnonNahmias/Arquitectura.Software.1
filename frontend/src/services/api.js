import axios from 'axios';

const API_URL = 'http://localhost:8080'; // Asegúrate de que esta URL sea correcta

export const fetchCourses = async () => {
  try {
    const response = await axios.get(`${API_URL}/courses`);
    return response.data;
  } catch (error) {
    console.error('Error fetching courses', error);
    throw error;
  }
};

export const loginUser = async (username, password) => {
  try {
    const response = await axios.post(`${API_URL}/login`, { username, password });
    return response.data; // Asume que la respuesta contiene el tipo de usuario
  } catch (error) {
    console.error('Error logging in', error);
    throw error;
  }
};

export const registerUser = async (username, password) => {
  const response = await fetch('/register', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ username, password }),
  });

  return response;
};