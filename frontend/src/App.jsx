import { Routes, Route } from 'react-router-dom'
import Home from './Home'
import Post from './Post'
import CreatePost from './CreatePost'
import Login from './Login'
import Navbar from './Navbar'


function App() {
  return (
    <>
      <Navbar />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/posts/:id" element={<Post />} />
        <Route path="/posts/newpost" element={<CreatePost />} />
        <Route path="/login" element={<Login />} />
      </Routes>
    </>
  )
}

export default App
