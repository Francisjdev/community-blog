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
  function handleSubmit(e) {
    e.preventDefault();
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
    <main className="post-page">
      <section className="post-form">

        <header className="post-header">
          <h1 className="post-title">New Post</h1>
        </header>

        <form onSubmit={handleSubmit}>

          <div className="form-group">
            <label htmlFor="title">Title</label>
            <input
              id="title"
              type="text"
              value={form.title}
              onChange={e =>
                setForm({...form, title: e.target.value})
              }
            />
          </div>

          <div className="form-group">
            <label htmlFor="slug">Slug</label>
            <input
              id="slug"
              type="text"
              value={form.slug}
              onChange={e =>
                setForm({...form, slug: e.target.value})
              }
            />
          </div>

          <div className="form-group">
            <label htmlFor="content">Content</label>
            <textarea
              id="content"
              rows="12"
              value={form.markdown_content}
              onChange={e =>
                setForm({...form, markdown_content: e.target.value})
              }
            />
          </div>

          <div className="form-group">
            <label htmlFor="image">Image URL</label>
            <input
              id="image"
              type="url"
              value={form.cover_image_url}
              onChange={e =>
                setForm({...form, cover_image_url: e.target.value})
              }
            />
          </div>

          <div className="form-footer">
            <button type="submit">
              Create Post
            </button>

            {message && (
              <p className="message">{message}</p>
            )}
          </div>

        </form>

      </section>
    </main>
  )
}

export default CreatePost
