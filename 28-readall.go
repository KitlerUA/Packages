
//START OMIT

switch {
case strings.HasPrefix(contentType, echo.MIMEApplicationJSON):
	rawBody, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Errorf("read body failed: %s", err)
		return response.ErrInternalError
	}

	log.Debugf("payload: %s", string(rawBody))

	err = json.Unmarshal(rawBody, &payload)
	if err != nil {
		log.Errorf("cannot unmarshal given payload: %s", err)
		return response.ErrBadRequest
	}

//END OMIT