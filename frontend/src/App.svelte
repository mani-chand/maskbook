<script>
  import "./app.css";
  import Topbar from "./components/Topbar.svelte";
  import Home from "./pages/Home.svelte";

  // 1. Import Router and Route from svelte-routing
  import { Router, Route } from "svelte-routing";

  // 2. Import your new 404 component (using the new name)
  import Page404 from "./pages/Page404.svelte";
  import Login from "./pages/Login.svelte";
  import Signup from "./pages/Signup.svelte";
  import Chat from "./pages/Chat.svelte";
  import Profile from "./pages/Profile.svelte";
</script>

<main>
  <!-- This Topbar will show on all pages -->
  <Topbar />
  <!-- Routing to pages -->
  <Router>
    <Route path="/">
      <Home />
    </Route>
    <Route path="/login">
      <Login />
    </Route>
    <Route path="/create">
      <Signup />
    </Route>
    <Route path="/chat">
      <Chat />
    </Route>
    <Route path="/profile/:user_id/">
      <Profile />
    </Route>
    <Route>
      <Page404 />
    </Route>
  </Router>
</main>
