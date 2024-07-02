import React, {useContext} from 'react'
import AuthContext from '../context/AuthContext'
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';

const LoginPage = () => {
    let {loginUser} = useContext(AuthContext)

    return (
        <Container>
            <Row className="justify-content-md-center">
                <Col xs lg="3" style={{ marginTop: '10%' }}>
                    <Form onSubmit={loginUser}>
                        <h3>Sign in</h3>
                        <Form.Group className="mb-3">
                            <Form.Label>Email address (admin@example.com)</Form.Label>
                            <Form.Control type="email" name="email" placeholder="admin@example.com" />
                        </Form.Group>

                        <Form.Group className="mb-3">
                            <Form.Label>Password (verysecret)</Form.Label>
                            <Form.Control id="passwordField" type="password" name="password" placeholder="Password" />
                        </Form.Group>
                        <Button variant="primary" type="submit">Submit</Button>
                    </Form>
                </Col>
            </Row>
        </Container>

    );
}
export default LoginPage