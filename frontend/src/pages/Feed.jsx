import { useState, useEffect } from 'react';
import { useAuth } from '../context/AuthContext';
import { tweetService } from '../services/api';
import TweetComposer from '../components/TweetComposer';
import Tweet from '../components/Tweet';
import Navbar from '../components/Navbar';
import '../App.css';

function Feed() {
  const [tweets, setTweets] = useState([]);
  const [loading, setLoading] = useState(true);
  const { user } = useAuth();

  useEffect(() => {
    loadTweets();
  }, []);

  const loadTweets = async () => {
    try {
      const response = await tweetService.getTweets();
      setTweets(response.data.tweets);
    } catch (error) {
      console.error('Failed to load tweets:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleTweetCreated = (newTweet) => {
    setTweets([newTweet, ...tweets]);
  };

  const handleTweetDeleted = (tweetId) => {
    setTweets(tweets.filter((tweet) => tweet.id !== tweetId));
  };

  if (loading) {
    return <div className="loading">Loading...</div>;
  }

  return (
    <div className="app">
      <Navbar />
      <div className="container">
        <h1>🐦 Twitter Clone</h1>
        <TweetComposer onTweetCreated={handleTweetCreated} />
        <div className="tweets">
          {tweets.length === 0 ? (
            <p className="no-tweets">No tweets yet. Be the first to tweet!</p>
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

export default Feed;
