package messages

// func Test_OnStart_ShouldSendIntroMessage(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	sender := mocks.NewMockMessageSender(controller)
// 	model := New(sender)

// 	sender.EXPECT().SendMessage("hello", int64(123))

// 	err := model.IncomingMessage(Message{
// 		Text:   "/start",
// 		UserID: 123,
// 	})

// 	assert.NoError(t, err)

// }

// func Test_UnknownCommand_ShouldAnswerWithHelpMessage(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	sender := mocks.NewMockMessageSender(controller)
// 	model := New(sender)

// 	sender.EXPECT().SendMessage("unknown command", int64(123))

// 	err := model.IncomingMessage(Message{
// 		Text:   "some command",
// 		UserID: 123,
// 	})

// 	assert.NoError(t, err)
// }
