require "lemmatizer"
require 'base64'
lem = Lemmatizer.new
attr_words = ARGV[2] ? Base64.decode64(ARGV[2]).split(' ') : []
str = ARGV[1] == 'base64' ? Base64.decode64(ARGV[0]): ARGV[0]
puts str.dup.split(' ').map { |x| attr_words.include?(x) ? x : lem.lemma(x) }.join(' ')
