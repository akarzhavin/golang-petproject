// import './App.css'
import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom'
import { AuthProvider } from './context/AuthContext';
import LoginPage from "./pages/LoginPage";
import HomePage from "./pages/HomePage";
import ArticlePage from "./pages/ArticlePage";


function App() {
  return (
    <div className="App">
      <main className="container-fluid">
        <BrowserRouter>
          <AuthProvider>
            <Routes>
                <Route element={<HomePage />} path='/' />
                <Route element={<ArticlePage />} path='/article/:articleId' />
                <Route element={<LoginPage />} path='/login' />
            </Routes>
          </AuthProvider>
        </BrowserRouter>
      </main>
    </div>
  )
}

export default App