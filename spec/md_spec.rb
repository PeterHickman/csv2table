# frozen_string_literal: true

require 'spec_helper'

describe 'markdown output' do
  before :all do
    build_program
  end

  after :all do
    remove_file('csv2')
  end

  after :each do
    remove_file('output.txt')
  end

  context 'from stdio' do
    context 'with the defaults' do
      it 'creates the markdown' do
        exec('cat ./spec/data.csv | ./csv2 --md > ./output.txt')

        contents_the_same('data.md')
      end
    end

    context 'delimiter is ,' do
      it 'creates the markdown' do
        exec('cat ./spec/data.csv | ./csv2 --md -delimit , > ./output.txt')

        contents_the_same('data.md')
      end
    end

    context 'delimiter is tab' do
      it 'creates the markdown' do
        exec('cat ./spec/data.tsv | ./csv2 --md -delimit "\t" > ./output.txt')

        contents_the_same('data.md')
      end
    end
  end

  context 'as argument' do
    context 'with the defaults' do
      it 'creates the markdown' do
        exec('./csv2 --md ./spec/data.csv > ./output.txt')

        contents_the_same('data.md')
      end
    end

    context 'delimiter is ,' do
      it 'creates the markdown' do
        exec('./csv2 -delimit , --md ./spec/data.csv > ./output.txt')

        contents_the_same('data.md')
      end
    end

    context 'delimiter is tab' do
      it 'creates the markdown' do
        exec('./csv2 -delimit "\t" --md ./spec/data.tsv > ./output.txt')

        contents_the_same('data.md')
      end
    end
  end
end
