import React, { useContext } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Button from '@mui/material/Button';
import Box from '@mui/material/Box';
import logo from '../assets/My-online-course.svg';
import { AuthContext } from '../context/AuthContext'; // Asegúrate de que la ruta sea correcta

const Header = () => {
  const { isAuthenticated, userType, logout } = useContext(AuthContext);
  const navigate = useNavigate();

  const handleLogout = async () => {
    await logout();
    navigate('/login');
  };

  return (
    <AppBar position="static">
      <Toolbar sx={{ position: 'relative', zIndex: 1 }}>
        <Box
          component="img"
          sx={{
            height: 200, // Increase the height
            position: 'absolute', // Position it absolutely
            top: '50%', // Center it vertically
            transform: 'translateY(-50%)', // Center it vertically
            marginRight: 2,
            zIndex: 0,
          }}
          alt="My Online Courses"
          src={logo}
        />
        <Box sx={{ flexGrow: 1 }} />
        <Button color="inherit" component={Link} to="/">Home</Button>
        {isAuthenticated && userType === 'admin' && (
          <Button color="inherit" component={Link} to="/manage">Manage</Button>
        )}
        {isAuthenticated && userType === 'normal' && (
          <Button color="inherit" component={Link} to="/my-courses">My Courses</Button>
        )}
        <Button color="inherit" component={Link} to="/about">About us</Button>
        {isAuthenticated ? (
          <Button color="inherit" onClick={handleLogout}>Logout</Button>
        ) : (
          <>
            <Button color="inherit" component={Link} to="/login">Login</Button>
            <Button color="inherit" component={Link} to="/register">Register</Button>
          </>
        )}
      </Toolbar>
    </AppBar>
  );
};

export default Header;
