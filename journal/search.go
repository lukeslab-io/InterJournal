package journal

// IterateJournal
//func (j *Journal) IterateJournal() {
//	var count int
//	for {
//		key, val, err := j.DB.Next()
//		if err == pogreb.ErrIterationDone {
//			break
//		}
//		if err != nil {
//			log.Fatal(err)
//		}
//		count++
//		s, err := strconv.ParseInt(string(key), 10, 64)
//		if err != nil {
//			panic(err)
//		}
//		var msg Message
//		json.Unmarshal(val, &msg)
//
//		t := time.Unix(s, 0)
//		fmt.Println(style.Render(fmt.Sprintf("[*] %s %s", t.Format(time.RFC822), msg.Content)))
//	}
//}
