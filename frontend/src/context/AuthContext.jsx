import {
  createContext,
  useContext,
  useState,
} from "react";

import api from "../api/axios";

const AuthContext = createContext(null);

export function AuthProvider({ children }) {
  const [user, setUser] = useState(() => {
    const storedUser =
      localStorage.getItem("user");

    if (!storedUser) {
      return null;
    }

    try {
      return JSON.parse(storedUser);
    } catch {
      localStorage.removeItem("user");
      return null;
    }
  });

  /*
   * LOGIN
   */
  const login = async (
    email,
    password
  ) => {

    const response = await api.post(
      "/login",
      {
        email,
        password,
      }
    );

    const {
      access_token,
      refresh_token,
      user,
    } = response.data;

    localStorage.setItem(
      "access_token",
      access_token
    );

    localStorage.setItem(
      "refresh_token",
      refresh_token
    );

    if (user) {

      localStorage.setItem(
        "user",
        JSON.stringify(user)
      );

      setUser(user);

    } else {

      /*
       * If backend doesn't return user
       * information, create a minimal
       * user object from the email.
       */

      const loggedInUser = {
        email,
      };

      localStorage.setItem(
        "user",
        JSON.stringify(loggedInUser)
      );

      setUser(loggedInUser);
    }

    return response.data;
  };

  /*
   * LOGOUT
   */
  const logout = () => {

    localStorage.removeItem(
      "access_token"
    );

    localStorage.removeItem(
      "refresh_token"
    );

    localStorage.removeItem(
      "user"
    );

    setUser(null);

    window.location.href =
      "/login";
  };

  const value = {
    user,
    login,
    logout,
  };

  return (
    <AuthContext.Provider
      value={value}
    >
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  return useContext(AuthContext);
}