import { useState } from 'react';
import { Link } from 'react-router-dom';
import { tweetService } from '../services/api';

function Tweet({ tweet, currentUser, onDeleted }) {
  const [liked, setLiked] = useState(
    currentUser && tweet.likes?.includes(currentUser.id)
  );
  const [likeCount, setLikeCount] = useState(tweet.likes?.length || 0);

  const handleLike = async () => {
    try {
      await tweetService.likeTweet(tweet.id);
      setLiked(!liked);
      setLikeCount(liked ? likeCount - 1 : likeCount + 1);
    } catch (error) {
      console.error('Failed to like tweet:', error);
    }
  };

  const handleDelete = async () => {
    if (!window.confirm('Are you sure you want to delete this tweet?')) return;

    try {
      await tweetService.deleteTweet(tweet.id);
      if (onDeleted) {
        onDeleted(tweet.id);
      }
    } catch (error) {
      console.error('Failed to delete tweet:', error);
    }
  };

  const isOwnTweet = currentUser && currentUser.username === tweet.username;

  return (
    <div className="tweet">
      <div className="tweet-header">
        <div>
          <Link to={`/profile/${tweet.username}`} className="tweet-author">
            <strong>@{tweet.username}</strong>
          </Link>
          <span className="tweet-date">
            {new Date(tweet.created_at).toLocaleDateString()}
          </span>
        </div>
        {isOwnTweet && (
          <button onClick={handleDelete} className="delete-btn">
            🗑️
          </button>
        )}
      </div>
      <p className="tweet-content">{tweet.content}</p>
      <div className="tweet-actions">
        <button
          onClick={handleLike}
          className={`like-btn ${liked ? 'liked' : ''}`}
          disabled={!currentUser}
        >
          {liked ? '❤️' : '🤍'} {likeCount}
        </button>
      </div>
    </div>
  );
}

export default Tweet;
