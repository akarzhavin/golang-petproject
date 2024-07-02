import React, {useLayoutEffect, useState} from 'react';
import {useAuthContext} from "../context/AuthContext";
import { useNavigate } from "react-router-dom";
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Header from "../components/Header";
import {useParams} from 'react-router-dom'

import {getApiHost} from "../settings";


const ArticlePage = () => {
    const {articleId} = useParams()
    let [article, setArticle] = useState({})
    const authContext = useAuthContext()
    const navigate = useNavigate()

    useLayoutEffect(() => {
        fetch(getApiHost() + `:8080/api/article/${articleId}`, {
            method:'GET',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json;charset=UTF-8',
                'Authorization': 'Bearer ' + authContext.authTokens.access
            },
        })
            .then(response => checkStatusCode202(response))
            .then(response => response.json())
            .then(data => setArticle(data.data))
            .catch(e => alert(e.message))

        return () => setArticle({})

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
                        <div style={{
                            padding: '10px',
                            marginTop: '10px',
                            marginBottom: '20px',
                        }}>
                            <a
                                href="#"
                                onClick={() => navigate('/')}
                                style={{
                                    color: 'white',
                                }}
                            >Back</a>
                            <div style={{
                                fontSize: '30px',
                                color: 'rgb(255 255 255 / 65%)',
                                fontWeight: 'bold',
                            }}>{article.title}</div>
                            <div>{article.created_at}</div>
                            <div style={{
                                marginTop: '10px',
                                marginBottom: '10px',
                                width: '100%',
                                height: '300px',
                                backgroundImage: `url(${article.image})`,
                                backgroundSize: 'cover',
                            }}></div>
                            <pre>{article.text}</pre>
                        </div>
                    </Col>
                </Row>
            </Container>
    );
};

export default ArticlePage;
