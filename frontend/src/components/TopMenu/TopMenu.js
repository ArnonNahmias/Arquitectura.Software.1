// TopMenu.js
import React from 'react';
import { Container, NavLink, Navbar, Nav } from 'react-bootstrap';
import { ReactComponent as Logo } from '../../assets/svg/LogoProvisorio.svg';
import './TopMenu.scss';

const TopMenu = ({ userRole, onSignOut }) => (
  <Navbar bg="dark" variant="dark" className="top-menu">
    <Container>
      <BrandNav />
      <MenuNav userRole={userRole} onSignOut={onSignOut} />
    </Container>
  </Navbar>
);

const BrandNav = () => (
  <Navbar.Brand>
    <Logo />
    <h2></h2>
  </Navbar.Brand>
);

const MenuNav = ({ userRole, onSignOut }) => (
  <Nav>
    {userRole && <NavLink href="#" onClick={onSignOut} className="sign-in-link">Sign out</NavLink>}
    {userRole === 'commonUser' && <NavLink href="#">Mis cursos</NavLink>}
  </Nav>
);

export default TopMenu;
