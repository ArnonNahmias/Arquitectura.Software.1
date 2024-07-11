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
  try {
    const response = await axios.post(`${API_URL}/register`, { username, password });
    return response.data;
  } catch (error) {
    console.error('Error registering', error);
    throw error;
  }
};

export const logoutUser = async () => {
  try {
    await axios.post(`${API_URL}/logout`);
  } catch (error) {
    console.error('Error logging out', error);
    throw error;
  }
};
