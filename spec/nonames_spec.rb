# frozen_string_literal: true

require 'spec_helper'

describe 'no column names in data' do
  before :all do
    build_program
  end

  after :all do
    remove_file('csv2')
  end

  after :each do
    remove_file('output.txt')
  end

  context 'table output' do
    it 'creates table output' do
      exec('./csv2 --nonames --names username,password,uid,gid,gecos,home,shell --delimit : --table ./spec/passwd > output.txt')

      contents_the_same('passwd.txt')
    end
  end

  context 'markdown output' do
    it 'creates markdown output' do
      exec('./csv2 --nonames --names username,password,uid,gid,gecos,home,shell --delimit : --md ./spec/passwd > output.txt')

      contents_the_same('passwd.md')
    end
  end

  context 'json output' do
    it 'creates json output' do
      exec('./csv2 --nonames --names username,password,uid,gid,gecos,home,shell --delimit : --json ./spec/passwd > output.txt')

      contents_the_same('passwd.json')
    end
  end
end
