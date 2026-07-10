import { useState } from 'react'
import { useNavigate } from 'react-router-dom'



function CreatePost() {
  const [form, setForm] = useState({
    title: "",
    slug: "",
    markdown_content: "",
    cover_image_url: "",
  })
  const [message, setMessage] = useState("")
  const navigate = useNavigate()
  function handleSubmit() {
    const token = localStorage.getItem('token')
    fetch('http://localhost:8080/api/posts/create', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`
      },
      body: JSON.stringify(form)
    })
      .then(res => res.json())
      .then(data => {console.log(data),setMessage("Post created successfully")
        setTimeout(() => navigate('/'), 2000)})

  }

  return (

    <div>
      <div><h1>New Post</h1></div>
      <label>Title</label>
      <input
        value={form.title}
        onChange={e => setForm({...form, title: e.target.value})}
      />
      <label>Slug</label>
      <input
        value={form.slug}
        onChange={e => setForm({...form, slug: e.target.value})}
      />
      <label>Content</label>
      <input
        value={form.markdown_content}
        onChange={e => setForm({...form, markdown_content: e.target.value})}
      />
      <label>Image url</label>
      <input
        value={form.cover_image_url}
        onChange={e => setForm({...form, cover_image_url: e.target.value})}
      />
      <button onClick={handleSubmit}>Create Post</button>
      <p>{message}</p>
    </div>
  )
}

export default CreatePost
