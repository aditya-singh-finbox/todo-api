import { useAuth } from "../context/AuthContext";

function Navbar() {
  const {
    user,
    logout,
  } = useAuth();

  return (
    <nav className="navbar">

      <div className="navbar-left">
        <h1>Todo App</h1>
      </div>

      <div className="navbar-right">

        <span>
          Welcome,{" "}
          {user?.name ||
            user?.email ||
            "User"}
        </span>

        <button
          className="logout-button"
          onClick={logout}
        >
          Logout
        </button>

      </div>

    </nav>
  );
}

export default Navbar;