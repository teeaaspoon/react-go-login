import React, {createContext, useContext, useState} from "react";

type AuthContextType = {
  user: any;
  setUser: React.Dispatch<any>;
}
const AuthContext = createContext<AuthContextType | null>(null);

// useAuth is a custom hook I made that allows any of the children components under AuthProvider to be
// able to access user and setUser.. see Login.tsx for how it is used.
export const useAuth = () => {
  const context = useContext(AuthContext);
  return [context?.user, context?.setUser];
}

// This is the provider that wraps around children components which need to access user and setUser
// we can achieve a similar functionality with redux but I wanted to try something new :)
const AuthProvider = ({children}: any) => {
  const [user, setUser] = useState(null);

  return (
    <AuthContext.Provider value={{user, setUser}}>
      {children}
    </AuthContext.Provider>
  )
}

export default AuthProvider;
