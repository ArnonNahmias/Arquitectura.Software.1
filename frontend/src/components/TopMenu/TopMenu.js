import React from 'react';
import { Container, NavLink, Navbar, Nav, Form, FormControl } from 'react-bootstrap';
import { ReactComponent as Logo } from '../../assets/svg/LogoProvisorio.svg';

import './TopMenu.scss';

const TopMenu = () => (
  <Navbar bg="dark" variant="dark" className="top-menu">
    <Container>
      <BrandNav />
      <MenuNav />
      <SearchBar />
    </Container>
  </Navbar>
);

const BrandNav = () => (
  <Navbar.Brand>
    <Logo />
    <h2></h2>
  </Navbar.Brand>
);

const MenuNav = () => (
  <Nav>
    <NavLink href="#" className="sign-in-link">Sign out</NavLink>
    <NavLink href="#">Mis cursos</NavLink>
  </Nav>
);

const SearchBar = () => (
  <Form inline>
    <FormControl type="text" placeholder="Search" className="mr-sm-2" />
    <button className="btn btn-outline-success my-2 my-sm-0" type="submit">Search</button>
  </Form>
);

export default TopMenu;