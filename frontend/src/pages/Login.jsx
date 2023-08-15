const Login = () => {
    return (
        <form className="p-1">
            <div class="mb-3">
                <label for="exampleInputEmail1" class="form-label">Email address</label>
                <input type="email" class="form-control" id="exampleInputEmail1" aria-describedby="emailHelp" />
                <div id="emailHelp" class="form-text">We'll never share your email with anyone else.</div>
            </div>
            <div class="mb-3">
                <label for="exampleInputPassword1" class="form-label">Password</label>
                <input type="password" class="form-control" id="exampleInputPassword1" />
            </div>
            <p>Belum punya akun? <a href="/register">Daftar di sini</a></p>
            <button type="submit" class="btn btn-primary">Login</button>
            <p>Tentang saya<a href="/about">About</a></p>
        </form>
    )
}

export default Login