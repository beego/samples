require "lemmatizer"
require 'base64'
lem = Lemmatizer.new
str = ARGV[1] == 'base64' ? Base64.decode64(ARGV[0]): ARGV[0]
puts str.dup.split(' ').map { |x| lem.lemma(x) }.join(' ')
