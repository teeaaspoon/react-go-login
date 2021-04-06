import React, { createContext, useContext, useState } from 'react';

type AuthContextType = {
  user: any;
  setUser: React.Dispatch<any>;
}

// Please note this is only exported for use in Login.test.tsx,
// If you need to useContext you should use useAuth() custom hook.
export const AuthContext = createContext<AuthContextType>({} as AuthContextType);

// useAuth is a custom hook I made that allows any of the children components under AuthProvider to be
// able to access user and setUser.. see Dashboard.tsx for how it is used.
export const useAuth = () => {
  const context = useContext(AuthContext);
  return {user: context?.user, setUser: context?.setUser as React.Dispatch<any>};
}

// This is the provider that wraps around children components which need to access user and setUser
// we can achieve a similar functionality with redux but I wanted to try something new :)
export const AuthProvider = ({children}: any) => {
  const [user, setUser] = useState(null);

  return (
    <AuthContext.Provider value={{user, setUser}}>
      {children}
    </AuthContext.Provider>
  )
}

