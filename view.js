function (doc, meta) { 
	if(doc.type === 'event') {
		emit(doc.likes, null);
	}
}