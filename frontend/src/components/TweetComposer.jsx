import { useState } from 'react';
import { tweetService } from '../services/api';

function TweetComposer({ onTweetCreated }) {
  const [content, setContent] = useState('');
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (!content.trim()) return;

    setError('');
    setLoading(true);

    try {
      const response = await tweetService.createTweet(content);
      setContent('');
      if (onTweetCreated) {
        onTweetCreated(response.data.tweet);
      }
    } catch (err) {
      setError(err.response?.data?.error || 'Failed to create tweet');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="tweet-composer">
      {error && <div className="error">{error}</div>}
      <form onSubmit={handleSubmit}>
        <textarea
          placeholder="What's happening?"
          value={content}
          onChange={(e) => setContent(e.target.value)}
          maxLength="280"
          rows="3"
        />
        <div className="composer-footer">
          <span className="char-count">{content.length}/280</span>
          <button type="submit" disabled={loading || !content.trim()}>
            {loading ? 'Tweeting...' : 'Tweet'}
          </button>
        </div>
      </form>
    </div>
  );
}

export default TweetComposer;
