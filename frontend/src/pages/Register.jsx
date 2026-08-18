import { useState } from "react";
import {
  Link,
  useNavigate,
} from "react-router-dom";

import api from "../api/axios";

function Register() {

  const navigate =
    useNavigate();

  const [name, setName] =
    useState("");

  const [email, setEmail] =
    useState("");

  const [password, setPassword] =
    useState("");

  const [error, setError] =
    useState("");

  const [loading, setLoading] =
    useState(false);

  const handleSubmit = async (e) => {

    e.preventDefault();

    setError("");

    if (
      !name ||
      !email ||
      !password
    ) {

      setError(
        "All fields are required"
      );

      return;
    }

    setLoading(true);

    try {

      await api.post(
        "/register",
        {
          name,
          email,
          password,
        }
      );

      navigate("/login");

    } catch (error) {

      setError(
        error.response?.data?.error ||
        "Registration failed"
      );

    } finally {

      setLoading(false);

    }
  };

  return (
    <div className="auth-container">

      <div className="auth-card">

        <h1>Todo App</h1>

        <h2>Create Account</h2>

        {error && (
          <div className="error">
            {error}
          </div>
        )}

        <form onSubmit={handleSubmit}>

          <div className="form-group">

            <label>
              Name
            </label>

            <input
              type="text"
              value={name}
              onChange={(e) =>
                setName(e.target.value)
              }
              placeholder="Enter your name"
            />

          </div>

          <div className="form-group">

            <label>
              Email
            </label>

            <input
              type="email"
              value={email}
              onChange={(e) =>
                setEmail(e.target.value)
              }
              placeholder="Enter your email"
            />

          </div>

          <div className="form-group">

            <label>
              Password
            </label>

            <input
              type="password"
              value={password}
              onChange={(e) =>
                setPassword(
                  e.target.value
                )
              }
              placeholder="Create a password"
            />

          </div>

          <button
            type="submit"
            disabled={loading}
          >
            {loading
              ? "Creating account..."
              : "Register"}
          </button>

        </form>

        <p className="auth-link">

          Already have an account?

          {" "}

          <Link to="/login">
            Login
          </Link>

        </p>

      </div>

    </div>
  );
}

export default Register;