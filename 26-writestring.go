package main

//START OMIT

case "id":
	id := req.Header.Get(echo.HeaderXRequestID)
	if id == "" {
		id = res.Header().Get(echo.HeaderXRequestID)
		if id == "" {
			return buf.WriteString(config.EmptyValuePlaceholder)
		}
	}
	return buf.WriteString(id)
case "remote_ip":
	ip := c.RealIP()
	if ip == "" {
		return buf.WriteString(config.EmptyValuePlaceholder)
	}
	return buf.WriteString(ip)

//END OMIT