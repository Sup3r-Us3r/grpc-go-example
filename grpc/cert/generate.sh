# Remove all .pem, .slr and .cnf files
rm -rf *.{pem,slr,cnf}

# /C=BR is for country
# /ST=Minas Gerais is for state or province
# /L=Belo Horizonte is for locality name or city
# /O=Example is for organisation
# /OU=Education is for organisation unit
# /CN=*.test is for common name or domain name
# /emailAddress=test@gmail.com is for email address

# 1. Generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=BR/ST=Minas Gerais/L=Belo Horizonte/O=Example/OU=Education/CN=*.test/emailAddress=test@gmail.com"

echo "CA's self-signed certificate"
openssl x509 -in ca-cert.pem -noout -text

# 2. Generate web server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/C=BR/ST=Minas Gerais/L=Belo Horizonte/O=Example/OU=Education/CN=*.test/emailAddress=test@gmail.com"

# 3. Use CA's private key to sign web server's CSR and get back the signed certificate
echo "subjectAltName=DNS:*.test.com,DNS:*.test.org,IP:0.0.0.0" > server-ext.cnf
openssl x509 -req -in server-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.cnf

echo "Server's signed certificate"
openssl x509 -in server-cert.pem -noout -text

# 4. Generate client's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout client-key.pem -out client-req.pem -subj "/C=BR/ST=Minas Gerais/L=Belo Horizonte/O=Example/OU=Education/CN=*.test/emailAddress=test@gmail.com"

# 5. Use CA's private key to sign client's CSR and get back the signed certificate
echo "subjectAltName=DNS:*.someclient.com,IP:0.0.0.0" > client-ext.cnf
openssl x509 -req -in client-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out client-cert.pem -extfile client-ext.cnf

echo "Client's signed certificate"
openssl x509 -in client-cert.pem -noout -text
