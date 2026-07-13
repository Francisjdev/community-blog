import { useState } from 'react'
import { useNavigate } from 'react-router-dom'


function Login() {
  const [form, setForm] = useState({
    email: "",
    password: "",
  })
  const navigate = useNavigate()

  function handleSubmit(e) {
    e.preventDefault(e)
    fetch('http://localhost:8080/api/auth/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(form)
    })
      .then(res => res.json())
      .then(data => {
        localStorage.setItem('token', data.access_token)
        navigate('/')
      })
  }
  return (
    <main className="auth-page">
      <section className="auth-form">

        <h1 className="auth-title">Login</h1>

        <form onSubmit={handleSubmit}>

          <div className="form-group">
            <label htmlFor="email">Email</label>
            <input
              id="email"
              type="email"
              value={form.email}
              onChange={e => setForm({...form, email: e.target.value})}
            />
          </div>

          <div className="form-group">
            <label htmlFor="password">Password</label>
            <input
              id="password"
              type="password"
              value={form.password}
              onChange={e => setForm({...form, password: e.target.value})}
            />
          </div>

          <button type="submit">
            Login
          </button>

        </form>

      </section>
    </main>
  )
}

export default Login
