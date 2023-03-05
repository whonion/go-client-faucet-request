for i in {1..490}
do
  output=$(nibid keys add AB$i)
  address=$(echo "$output" | grep -oP "(?<=address: )[a-z0-9]+")
  echo "$address" >> addresses.txt
done