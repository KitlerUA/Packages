
//START OMIT

	uploadPath := filepath.Join(h.conf.RootPath, data.Path)
	err = os.MkdirAll(filepath.Dir(uploadPath), h.conf.DirMode)
	if err != nil {
		log.Errorf("mkdirall failed: %s", err)
		return response.ErrInternalError
	}

	//END OMIT