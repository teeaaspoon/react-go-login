import React, {SyntheticEvent, useState} from 'react';
import { useHistory } from 'react-router-dom';
import {useAuth} from "../contexts/AuthContext";
import {backendService} from "../service/backendService";


const Login = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [user, setUser] = useAuth();
  const [errorMsg, setErrorMsg] = useState('');
  const history = useHistory();

  const onSubmit = async (e: SyntheticEvent) => {
    e.preventDefault();

    try {
      const response = await backendService.login(email, password);
      setErrorMsg('');
      localStorage.setItem("CSRFToken", response.headers['csrf-token']);
      setUser(response.data);
      // redirect to dashboard
      history.push({
        pathname: "/",
      });
    } catch (e) {
      // console.log(Object.keys(e), e.message);
      setErrorMsg(e.response?.data.error);
    }
  }

  return (
    <div>
      <form className="login-form" onSubmit={onSubmit}>
        <h1>Sign Into Your Account</h1>
        <div>
          <label htmlFor="email">Email Address</label>
          <input
            type="email"
            id="email"
            className="field"
            required
            onChange={e => setEmail(e.target.value)}
          />
        </div>

        <div>
          <label htmlFor="password">Password</label>
          <input
            type="password"
            id="password"
            className="field"
            required
            onChange={e => setPassword(e.target.value)}
          />
        </div>
        <input type="submit" value="Login to my Dashboard" className="button block"/>
      </form>
      <p>{errorMsg}</p>
    </div>
  )
}

export default Login;
