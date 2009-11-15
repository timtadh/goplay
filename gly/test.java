public class Heap
{
    int max_size;
    int heap[];
    int p = 1;
    
    public Heap(int max)
    {
        max_size = max;
        heap = new int[max_size+1];
    }
    
    public int getSize()
    {
        return p-1;
    }
    
    public int getMax()
    {
        return heap[1];
    }
    
    public int popMax()
    {
        int t = heap[1];
        remove(t);
        return t;
    }
    
    private void swap(int i, int j)
    {
        int t = heap[i];
        heap[i] = heap[j];
        heap[j] = t;
    }
    
    private void fixUp(int i)
    {
        //i/2 == parent of i
        for (; i > 1 && heap[i/2] < heap[i]; i = i/2)
        {
            swap(i, i/2);
        }
    }
    
    public void add(int x)
    {
        heap[p] = x;
        fixUp(p);
        p++;
    }
    
    private void fixDown(int i)
    {
        while(2*i <= p)
        {
            int j = 2*i;
            if (j < (p) && heap[j] < heap[j+1]) { j++; }
            if (heap[i] >= heap[j]) { break; }
            swap(i, j);
            i = j;
        }
    }
    
    private boolean rh(int x, int i)
    {
        
        if (i < p)
        {
            if (x == heap[i])
            {
                if (i + 1 == p)
                {
                    heap[i] = 0;
                    p--;
                }
                else
                {/*
                    System.out.print(">> ");
                    System.out.print(heap[i]);
                    System.out.print(", ");
                    System.out.println(heap[p-1]);*/
                    heap[i] = heap[p-1];
                    p--;
                    fixDown(i);
                }
                return true;
            }
            boolean v = false;
            if (x <= heap[2*i]) //left child
            {
                v = rh(x, 2*i);
            }
            if (v == false && x <= heap[2*i + 1]) //right child
            {
                v = rh(x, 2*i+1);
            }
            return v;
        }
        else
        {
            return false;
        }
    }
    
    public void remove(int x)
    {
        int i = 1;
        rh(x, i);
    }
    
    public void print()
    {
        for (int i = 1; i < p; i++)
        {
            System.out.print(heap[i]);
            System.out.print(", ");
        }
        System.out.println();
    }
}

