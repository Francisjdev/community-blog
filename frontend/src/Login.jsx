import { useState } from 'react'
import { useNavigate } from 'react-router-dom'


function Login() {
  const [form, setForm] = useState({
    email: "",
    password: "",
  })
  const navigate = useNavigate()

  function handleSubmit() {
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
    <div>
      <h1>Login</h1>
      <label>Email</label>
      <input
        value={form.email}
        onChange={e => setForm({...form, email: e.target.value})}
      />
      <label>Password</label>
      <input
        type="password"
        value={form.password}
        onChange={e => setForm({...form, password: e.target.value})}
      />
      <button onClick={handleSubmit}>Login</button>
    </div>
  )
}

export default Login
