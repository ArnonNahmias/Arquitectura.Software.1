import React, { createContext, useState, useEffect } from 'react';
import { loginUser, logoutUser } from '../services/api'; // Asegúrate de que la ruta sea correcta

export const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [userType, setUserType] = useState('');

  const login = async (username, password) => {
    try {
      const user = await loginUser(username, password);
      setIsAuthenticated(true);
      setUserType(user.type); // Asume que la respuesta del login contiene el tipo de usuario
    } catch (error) {
      throw error;
    }
  };

  const logout = async () => {
    try {
      await logoutUser();
      setIsAuthenticated(false);
      setUserType('');
    } catch (error) {
      throw error;
    }
  };

  useEffect(() => {
    // Verificar si el usuario ya está autenticado (por ejemplo, verificando la cookie de token)
    // Aquí deberías agregar la lógica para verificar si el usuario está autenticado
  }, []);

  return (
    <AuthContext.Provider value={{ isAuthenticated, userType, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};
