authors=('ChatGPT4' 'Copilot' 'Gemini2.5')

rm times.txt
for author in "${authors[@]}"
do
  cd $author
  for file in *.java
	do
    ./build.sh $file
		echo $author $file >> ../times.txt
		for cnt in {1..7}
		do
			{ time ./run.sh $file ; } 2>> ../times.txt
		done
	done
  cd ..
done