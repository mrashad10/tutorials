<?php
//mail service start
include("class.phpmailer.php");
include("class.smtp.php");
$mail = new PHPMailer();
$mail->IsSMTP(); // set mailer to use SMTP
$mail->Mailer = "smtp";
$mail->Host = getenv('SMTP_HOST'); // specify main and backup server // Twinbrothers SMTP server Name "gains.lathalinuxcloud.com";
$mail->Port = 465; // set the port to use
$mail->SMTPAuth = true; // turn on SMTP authentication
$mail->Username = getenv('SMTP_USERNAME'); // your SMTP username or your gmail username
$mail->Password = getenv('SMTP_PASSWORD'); // give your gmail password
$mail->SMTPDebug = 0;
$mail->SMTPSecure = 'ssl';

$from1 = getenv('SMTP_FROM'); // Reply to this email
