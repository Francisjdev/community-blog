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
    <div>
      <div>Post page</div>
      <h2>{post.title}</h2>
      <h3>  { post.slug}</h3>
      <p>{post.markdown_content}</p>

      <button onClick={handleDelete}>Delete Post</button>
      <p>{message}</p>
    </div>
  )
}

export default Post
