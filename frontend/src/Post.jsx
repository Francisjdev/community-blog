import { useState, useEffect } from 'react'
import { useParams } from 'react-router-dom'
import { useNavigate } from 'react-router-dom'

function Post() {
  const [post, setPosts] = useState({})
  const { id } = useParams()
  const navigate = useNavigate()
  const [message, setMessage] = useState("")
  useEffect(() => {
    fetch(`http://localhost:8080/api/posts/${id}`)
      .then(res => res.json())
      .then(data => setPosts(data))
  }, [id])
  function handleDelete() {
    const token = localStorage.getItem('token')
    fetch('http://localhost:8080/api/posts/deletesinglepost', {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`
      },
      body: JSON.stringify({ post_id: id })
    })
      .then(res => res.json())
      .then(data => {console.log(data),setMessage("Post deleted successfully")
        setTimeout(() => navigate('/'), 2000)})

  }
   return (
    <main className="post-page">
      <div className="container">
        <article className="post">

          <header className="post-header">
            <h1 className="post-title">{post.title}</h1>

            <div className="post-meta">
              {post.author && <span>By {post.author}</span>}
              {post.created_at && (
                <>
                  {" • "}
                  {new Date(post.created_at).toLocaleDateString()}
                </>
              )}
            </div>

            <p className="post-slug">/{post.slug}</p>
          </header>

          <section className="post-content">
            <p>{post.markdown_content}</p>
          </section>

          <footer className="post-footer">
            <button
              className="delete-btn"
              onClick={handleDelete}
            >
              Delete Post
            </button>

            {message && (
              <p className="message">{message}</p>
            )}
          </footer>

        </article>
      </div>
    </main>
  );
}

export default Post
