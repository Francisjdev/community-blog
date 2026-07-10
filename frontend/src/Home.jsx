import { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
import placeholder from './assets/placeholder.jpg'

const PLACEHOLDER = placeholder

function Home() {
  const [posts, setPosts] = useState([])

  useEffect(() => {
    fetch('http://localhost:8080/api/posts')
      .then(res => res.json())
      .then(data => setPosts(data))
  }, [])

  const featured = posts[0]
  const rest = posts.slice(1)
console.log(featured?.cover_image_url, PLACEHOLDER)
  return (
    <div className="home">
      {featured && (
        <Link to={`/posts/${featured.post_id}`} className="featured">
          <img src={featured.cover_image_url || PLACEHOLDER} alt={featured.title} />
          <div className="featured-text">
            <span className="featured-slug">{featured.slug}</span>
            <h1>{featured.title}</h1>
            <p>{featured.meta_description}</p>
          </div>
        </Link>
      )}
      <div className="post-grid">
        {rest.map(post => (
          <Link to={`/posts/${post.post_id}`} className="post-card" key={post.post_id}>
            <img src={post.cover_image_url || PLACEHOLDER} alt={post.title} />
            <span className="post-slug">{post.slug}</span>
            <h3>{post.title}</h3>
            <p>{post.meta_description}</p>
          </Link>
        ))}
      </div>
    </div>
  )
}

export default Home
