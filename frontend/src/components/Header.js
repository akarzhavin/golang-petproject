// App.js
import React from 'react';
import {useAuthContext} from "../context/AuthContext";
import Container from 'react-bootstrap/Container';
import Navbar from 'react-bootstrap/Navbar';
import NavDropdown from 'react-bootstrap/NavDropdown';
import './Header.css'
import { useNavigate } from "react-router-dom";


const Header = () => {
    const authContext = useAuthContext()
    const navigate = useNavigate()

    return (
        <Navbar className="blog-header">
            <Container>
                <Navbar.Brand href="#" onClick={() => navigate("/")} style={{color: 'white'}}>Home blog</Navbar.Brand>
                <Navbar.Toggle />
                <Navbar.Collapse className="justify-content-end">
                    <Navbar.Text style={{color: 'rgb(255 255 255 / 65%)'}}>
                        Signed in as:  <NavDropdown title={authContext.userData?.email} id="collapsible-nav-dropdown">
                        <NavDropdown.Item href="#action4">
                            <div href="#logout" onClick={() => authContext.logoutUser()}>Logout</div>
                        </NavDropdown.Item>
                    </NavDropdown>
                    </Navbar.Text>
                </Navbar.Collapse>
            </Container>
        </Navbar>
    );
};

export default Header;
