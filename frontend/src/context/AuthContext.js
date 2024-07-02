import { createContext, useState, useEffect, useContext } from "react";
import jwt_decode from 'jwt-decode'
import { useNavigate } from "react-router-dom";
import { getApiHost } from "../settings";
const AuthContext = createContext(null)
const accessTokensLocalName = 'accessTokens'


export default AuthContext;

const extractTokensFromLocalStorage = () => {
    let accessTokens, userData;

    if (localStorage.getItem(accessTokensLocalName)) {
        accessTokens = JSON.parse(localStorage.getItem(accessTokensLocalName));
        userData = jwt_decode(accessTokens.access);
    } else {
        accessTokens = null;
        userData = null;
    }

    return {
        accessTokens: accessTokens, 
        userData: userData
    }
}

export const AuthProvider = ({children}) => {
    let [authTokens, setAuthTokens] = useState(() => extractTokensFromLocalStorage().accessTokens)
    let [user, setUser] = useState(() => extractTokensFromLocalStorage().userData)
    let [loading, setLoading] = useState(true)

    const navigate = useNavigate()

    const saveTokens = (data) => {
        setAuthTokens(data)
        setUser(jwt_decode(data.access))

        localStorage.setItem(accessTokensLocalName, JSON.stringify(data))
    }

    const checkStatusCode202 = (response) => {
        if (response.status === 202) {
            return response
        } else if (response.status === 401) {
            throw Error('Unauthorized!')
        }

        throw Error('Something went wrong!')
    }

    let loginUser = (e) => {
        e.preventDefault()

        fetch(getApiHost() + ':8081/api/auth/token/', {
            method:'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json;charset=UTF-8'    
            },
            body:JSON.stringify({
                'email':e.target.email.value,
                'password':e.target.password.value
            })
        })
        .then((response) => checkStatusCode202(response))
        .then((response) => response.json())
        .then((data) => saveTokens(data.data))
        .then(() => navigate('/'))
        .catch(e => alert(e.message))
    }

    const logoutUser = () => {
        setAuthTokens(null);
        setUser(null);
        localStorage.removeItem(accessTokensLocalName)
        navigate('/login')
    }

    let updateTokens = async(e) => {
        if (!authTokens?.refresh) {
            logoutUser()
            setLoading(false)
            return
        }

        let response = await fetch(getApiHost() + ':8081/api/auth/token/refresh/', {
            method:'POST',
            headers:{
                'Accept': 'application/json',
                'Content-Type': 'application/json;charset=UTF-8'    
            },
            body:JSON.stringify({'refresh': authTokens?.refresh})
        })

        if (response.status === 202){
            let data = await response.json()
            saveTokens(data.data)
        } else {
            logoutUser()
        }

        if(loading){
            setLoading(false)
        }
    }

    let constextData = {
        userData: user,
        authTokens: authTokens,
        loginUser: loginUser,
        logoutUser: logoutUser,
    }

    useEffect(() => {
        if (loading) {
            updateTokens()
        }

        let fourMinutes = 1000 * 60 * 4;

        let interval = setInterval(() => {
            if(authTokens) {
                updateTokens()
            }
        }, fourMinutes)
        return () => clearInterval(interval)

    }, [authTokens, loading])

    return (
        <AuthContext.Provider value={constextData}>
            {loading ? null : children}
        </AuthContext.Provider>
    )
}

export function useAuthContext() {
    return useContext(AuthContext)
}