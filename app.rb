require 'sinatra'

post '/file_upload' do
  redirect '/redirect', 302
end

get '/redirect' do
  return 'abcd'
end
