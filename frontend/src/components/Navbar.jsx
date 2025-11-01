import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';

function Navbar() {
  const { user, logout } = useAuth();
  const navigate = useNavigate();

  const handleLogout = () => {
    logout();
    navigate('/login');
  };

  return (
    <nav className="navbar">
      <div className="nav-content">
        <Link to="/" className="nav-brand">
          🐦 Twitter Clone
        </Link>
        <div className="nav-links">
          {user && (
            <>
              <Link to={`/profile/${user.username}`}>Profile</Link>
              <span className="nav-user">@{user.username}</span>
              <button onClick={handleLogout} className="logout-btn">
                Logout
              </button>
            </>
          )}
        </div>
      </div>
    </nav>
  );
}

export default Navbar;
