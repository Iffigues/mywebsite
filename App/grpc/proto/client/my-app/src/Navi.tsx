import React, { useState, useRef, useLayoutEffect}  from 'react';
import { Navbar, Container, Nav, Form, FormControl,Button } from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import Image from 'react-bootstrap/Image'
import Dropdown from 'react-bootstrap/Dropdown';

const Lin = React.forwardRef((props, ref: React.Ref<HTMLDivElement>) => (
  <Dropdown>
      <Dropdown.Toggle variant="success" id="dropdown-basic">
        Dropdown Button
      </Dropdown.Toggle>		 
<Dropdown.Menu>
        <Dropdown.Item href="/fortune">Fortune</Dropdown.Item>
        <Dropdown.Item href="/rig">Rig</Dropdown.Item>
        <Dropdown.Item href="/toiler">Toilet</Dropdown.Item>
	<Dropdown.Item href="/figlet">Figlet</Dropdown.Item>
	<Dropdown.Item href="/cowsay">Cowsay</Dropdown.Item>
      </Dropdown.Menu>
    </Dropdown>
));


function Navi() {
return (
	
<Navbar bg="light" expand="lg">
  <Container fluid>
    <Navbar.Toggle aria-controls="navbarScroll" />
    <Navbar.Collapse id="navbarScroll">
      <Nav
        className="me-auto my-2 my-lg-0"
        style={{ maxHeight: '100px' }}
        navbarScroll
      >
        <Nav.Link href="/">Home</Nav.Link>
        <Lin />
      </Nav>
    </Navbar.Collapse>
  </Container>
</Navbar>
  );
}

export default Navi
