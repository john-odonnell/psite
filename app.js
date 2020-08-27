// require external packages
const express = require("express");
const bodyParser = require("body-parser");
const ejs = require("ejs");
const _ = require("lodash");
const mongoose = require("mongoose");

// initialize express, use bodyParser and ejs
const app = express();
app.use(bodyParser.urlencoded( {extended: true} ));
app.set("view engine", "ejs");
// use public dir as static elements dir
app.use(express.static("public"));





// MONGO INIT
// connect to local database
mongoose.connect('mongodb://localhost:27017/blogdb',
{useNewUrlParser: true, useUnifiedTopology: true});
// used to escape depreciation warnings
// mongoose.set("useFindAndModify", false);
// create document schema and collection model
const postSchema = new mongoose.Schema({
  title: {
    type: String,
    required: [true, "Post requires a Title."]
  },
  body: {
    type: String,
    required: [true, "Post requires a Body."]
  },
  date: {
    type: Date,
    required: [true, "Post requires Date and Time data."]
  }
});
const Post = mongoose.model("Post", postSchema);





// GET REQUESTS
app.get("/", (req, res) => {
  // renders the static index of the server at root
  res.render(__dirname + "/views/index.ejs");
});

app.get("/blog", (req, res) => {
  // renders to the user the blog page
  // if there are no blog posts, fill the blog with the default lorem ipsum
  Post.find({}, (err, docs) => {
    res.render(__dirname + "/views/blog.ejs", {posts: docs});
  });
});

app.get("/blog/newpost", (req, res) => {
  // renders the page to add a new blog post to the db
  res.render(__dirname + "/views/newpost.ejs");
});

app.get("/blog/:title", (req, res) => {
  // renders the post page filled with db content pertaining to the
  // title passed to the server as url parameter
  let title = req.params.title;
  Post.findOne({title: title}, (err, doc) => {
    res.render(__dirname + "/views/post.ejs", {post: doc});
  });
});





// DISABLED UNTIL ADMIN AUTHENTICATION IS ADDED
// POST REQUEST
app.post("/newpost", (req, res) => {
  // post request passes post title and content to the server
  // the server forms the post into a json and adds to the db
  // redirects the user to the blog page

  // let title = req.body.title;
  // let body = req.body.content;
  // let date = new Date();
  //
  // let newPost = new Post({
  //   title: title,
  //   body: body,
  //   date: date
  // });
  // newPost.save();

  res.redirect("/blog");
});





// SERVER SPINUP
app.listen(3366, () => {
  // initialize the server on local host
  console.log("Server running on Local 3366");
});
