import { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';
import { userService, tweetService } from '../services/api';
import Tweet from '../components/Tweet';
import Navbar from '../components/Navbar';
import '../App.css';

function Profile() {
  const { username } = useParams();
  const [profile, setProfile] = useState(null);
  const [tweets, setTweets] = useState([]);
  const [loading, setLoading] = useState(true);
  const { user } = useAuth();

  useEffect(() => {
    loadProfile();
    loadTweets();
  }, [username]);

  const loadProfile = async () => {
    try {
      const response = await userService.getUserProfile(username);
      setProfile(response.data.user);
    } catch (error) {
      console.error('Failed to load profile:', error);
    } finally {
      setLoading(false);
    }
  };

  const loadTweets = async () => {
    try {
      const response = await tweetService.getUserTweets(username);
      setTweets(response.data.tweets);
    } catch (error) {
      console.error('Failed to load tweets:', error);
    }
  };

  const handleFollow = async () => {
    try {
      await userService.followUser(username);
      loadProfile();
    } catch (error) {
      console.error('Failed to follow user:', error);
    }
  };

  const handleTweetDeleted = (tweetId) => {
    setTweets(tweets.filter((tweet) => tweet.id !== tweetId));
  };

  if (loading) {
    return <div className="loading">Loading...</div>;
  }

  if (!profile) {
    return <div className="error">User not found</div>;
  }

  const isOwnProfile = user && user.username === username;
  const isFollowing = user && profile.followers && profile.followers.some(
    (followerId) => followerId === user.id
  );

  return (
    <div className="app">
      <Navbar />
      <div className="container">
        <div className="profile-header">
          <div className="profile-info">
            <h1>{profile.name || profile.username}</h1>
            <p className="username">@{profile.username}</p>
            {profile.bio && <p className="bio">{profile.bio}</p>}
            <div className="profile-stats">
              <span>{profile.followers?.length || 0} Followers</span>
              <span>{profile.following?.length || 0} Following</span>
            </div>
          </div>
          {!isOwnProfile && user && (
            <button onClick={handleFollow} className="follow-button">
              {isFollowing ? 'Unfollow' : 'Follow'}
            </button>
          )}
        </div>
        <div className="tweets">
          <h2>Tweets</h2>
          {tweets.length === 0 ? (
            <p className="no-tweets">No tweets yet</p>
          ) : (
            tweets.map((tweet) => (
              <Tweet
                key={tweet.id}
                tweet={tweet}
                currentUser={user}
                onDeleted={handleTweetDeleted}
              />
            ))
          )}
        </div>
      </div>
    </div>
  );
}

export default Profile;
