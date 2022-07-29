SELECT ID,USER,HOST,DB,COMMAND,TIME,STATE,INFO FROM information_schema.processlist WHERE command !='Sleep';
