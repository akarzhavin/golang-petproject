import { useNavigate } from "react-router-dom";
import React, {useLayoutEffect, useState} from 'react';
import {useAuthContext} from "../context/AuthContext";
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Header from "../components/Header";

import {getApiHost} from "../settings";


const HomePage = () => {
    const navigate = useNavigate()
    let [articles, setArticles] = useState([{}])
    const authContext = useAuthContext()

    useLayoutEffect(() => {
        fetch(getApiHost() + ':8080/api/article/list', {
            method:'GET',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json;charset=UTF-8',
                'Authorization': 'Bearer ' + authContext.authTokens.access
            },
        })
            .then(response => checkStatusCode202(response))
            .then(response => response.json())
            .then(data => setArticles(data.data))
            .catch(e => alert(e.message))

        return () => setArticles([{}])

    }, [])

    const checkStatusCode202 = (response) => {
        if (response.status === 202) {
            return response
        } else if (response.status === 401) {
            throw Error('Unauthorized!')
        }

        throw Error('Something went wrong!')
    }

    return (
            <Container>
                <br/>
                <Row>
                    <Header/>
                </Row>
                <br/>
                <Row>
                    <Col>
                        {articles.map((article, index) => (
                            <div key={index} style={{
                                padding: '10px',
                                marginTop: '10px',
                                marginBottom: '20px',
                            }}>
                                <a href="#" onClick={() => navigate(`/article/${article.id}`)} style={{
                                    fontSize: '30px',
                                    color: 'white',
                                    fontWeight: 'bold',
                                }}>{article.title}</a>
                                <div>{article.created_at}</div>
                                <div style={{
                                    marginTop: '10px',
                                    marginBottom: '10px',
                                    width: '100%',
                                    height: '300px',
                                    backgroundImage: `url(${article.image})`,
                                    backgroundSize: 'cover',
                                }}></div>
                                <p>{article.text}</p>
                                <a href="#" onClick={() => navigate(`/article/${article.id}`)} style={{
                                    color: 'white',
                                }}>Read full story</a>
                            </div>
                        ))}
                    </Col>
                </Row>
            </Container>
    );
};

export default HomePage;
