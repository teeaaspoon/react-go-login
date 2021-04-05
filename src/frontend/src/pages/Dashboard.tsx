import React, { useEffect, useState } from 'react';
import { backendService } from '../service/backendService';
import { useAuth } from '../contexts/AuthContext';
import { useHistory } from 'react-router-dom';

const Dashboard = () => {
  const [isLoading, setIsLoading] = useState(true);
  const [user, setUser] = useAuth();
  const [devices, setDevices] = useState(0);
  const history = useHistory();

  // try to make a request to /dashboard endpoint.
  // if it fails it means we're unauthenticated so redirect the user back to /login
  useEffect(() => {
    async function fetchDashboard() {
      try {
        // if the user has a JWT and a csrfToken, but reloads the page
        // the logout button will disappear because it is dependent on the state of the user
        // we need to do a service call to get the information of the user
        if (!user && localStorage.getItem("CSRFToken") !== null) {
          const userResponse = await backendService.getUser();
          setUser(userResponse.data);
        }
        const response = await backendService.getDashboard();
        setDevices(response.data.Devices);
        setIsLoading(false);
      } catch(e) {
        // if we get an error it means session timedout or they tried to access this link unauthenticated
        // setUser to null and localStorage cleared
        setUser(null);
        localStorage.clear();
        // redirect to login
        history.push({
          pathname: "/login",
        });
      }
    }

    fetchDashboard();
  }, []);

  return (
    <div>
      {isLoading ? (
        <p>Loading</p>
      ) : (
        <p>{devices} Devices Connected</p>
      )}

    </div>
  )
}

export default Dashboard;
