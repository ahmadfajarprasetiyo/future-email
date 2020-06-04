const express = require('express')
const app = express()
const port = 8081
const nodemailer = require("nodemailer")

var bodyParser = require('body-parser')

const emailSender = ''
const passwordSender = ''

async function sendMail(email, content) {
    // Generate test SMTP service account from ethereal.email
    let testAccount = await nodemailer.createTestAccount();
  
    // create reusable transporter object using the default SMTP transport
    let transporter = nodemailer.createTransport({
      host: "smtp.gmail.com",
      port: 587,
      secure: false, 
      auth: {
        user: emailSender, 
        pass: passwordSender, 
      },
    });

    // send mail with defined transport object
    let info = await transporter.sendMail({
      from: '"Ahmad Fajar" <'+emailSender+'>', 
      to: email, 
      subject: "Future Email", 
      text: content, 
    });
  
    console.log("Message sent: %s", info.messageId);
  }

function mailHandler(req, res) {
    var timeSend = req.body.time_send*1000
    var timeNow = new Date().getTime()
    var timeOut = timeSend-timeNow

    var email = req.body.email
    var content = req.body.content

    setTimeout(function(){ 
        sendMail(email, content)
    }, timeOut);
    
    res.send('Succes')
}

app.use(bodyParser.json())
app.post('/send', mailHandler)

app.listen(port, () => console.log(`Example app listening at http://localhost:${port}`))