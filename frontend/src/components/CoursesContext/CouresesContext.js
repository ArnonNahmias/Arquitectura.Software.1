//permite compartir los datos de cursos existentes entre admin y user
// CoursesContext.js
import React, { createContext, useState } from 'react';

export const CoursesContext = createContext();

export const CoursesProvider = ({ children }) => {
  const [courses, setCourses] = useState([]);

  return (
    <CoursesContext.Provider value={{ courses, setCourses }}>
      {children}
    </CoursesContext.Provider>
  );
};
