const express = require("express");
const axios = require("axios");
const app = express()
app.get("/", (req, res) => {
    res.send(axios.post("http://localhost:8081/api/v1/users"));
});
app.listen(3000, ()=>{
    console.log('server running at 3000')
})