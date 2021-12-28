```
Usage: tgsum [-bam bamName] [-sample sample] [-output outputDir] [-threads threads]
       统计三代下机数据bam的长度信息.

Options:
  -bam value
        Input bam (empty for stdin), for example: -bam=a.bam -bam=b.bam
  -help
        Display help
  -output string
        Output directory (default ".")
  -sample string
        Sample name (default "sample")
  -threads int
        Number of threads to use (0 = auto)

Example: tgsum -bam reads.bam    # 输出到当前目录sample_stat.{json,txt}
         tgsum -bam reads.bam -sample wangxuehan -output result    # 输出到result目录下wangxuehan_stat.{json,txt}
         tgsum -bam reads1.bam -bam reads2.bam    # 支持多个bam合并统计, 适用于一个样本多个数据

Result:
        header:Sample    Reads Number    Bases(bp)    Mean Length(bp)    Longest(bp)    N50(bp)

Author:     王雪涵
version:    v0.1
```
